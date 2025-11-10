package deltadefi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/sidan-lab/rum/wallet"
	"golang.org/x/crypto/pbkdf2"
)

// LoadOperationKey loads and decrypts the operation key required for transaction signing.
// This method must be called before performing any transaction operations like placing orders.
//
// Parameters:
//   - passcode: The operation passcode for decrypting the key
//
// Returns:
//   - error: nil on success, error on failure
func (d *DeltaDeFi) LoadOperationKey(passcode string) error {
	res, err := d.Accounts.GetOperationKey()
	if err != nil {
		return err
	}

	// Use proper AES-GCM decryption that matches frontend encryption format
	operationKey, err := decryptWithCipher(res.EncryptedOperationKey, passcode)
	if err != nil {
		return fmt.Errorf("decryption failed: %w", err)
	}

	operationWallet, err := wallet.NewRootKeyWallet(operationKey, wallet.NewDerivationIndices())
	if err != nil {
		return err
	}

	d.OperationWallet = operationWallet
	return nil
}

func decryptWithCipher(encryptedDataJSON string, password string) (string, error) {
	// Parse the encrypted data JSON
	var encData struct {
		IV         string  `json:"iv"`
		Salt       *string `json:"salt,omitempty"`
		Ciphertext string  `json:"ciphertext"`
	}

	err := json.Unmarshal([]byte(encryptedDataJSON), &encData)
	if err != nil {
		return "", fmt.Errorf("failed to parse encrypted data: %w", err)
	}

	// Decode IV from base64
	iv, err := base64.StdEncoding.DecodeString(encData.IV)
	if err != nil {
		return "", fmt.Errorf("failed to decode IV: %w", err)
	}

	// Decode ciphertext from base64
	ciphertext, err := base64.StdEncoding.DecodeString(encData.Ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// Handle salt - support both new format (with salt) and legacy format (without salt)
	var salt []byte
	if encData.Salt != nil && *encData.Salt != "" {
		// New format: use the provided salt
		salt, err = base64.StdEncoding.DecodeString(*encData.Salt)
		if err != nil {
			return "", fmt.Errorf("failed to decode salt: %w", err)
		}
	} else {
		// Legacy format: use zero-filled salt of IV length for backward compatibility
		salt = make([]byte, len(iv))
	}

	// Derive cryptographic key from password using PBKDF2
	// Matches frontend: 100,000 iterations, SHA-256, 256-bit key
	derivedKey := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// Create AES cipher
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Decrypt the data
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt (incorrect password or corrupted data): %w", err)
	}

	// Return the decrypted data as string
	return string(plaintext), nil
}

// PostOrder is a high-level method for placing an order.
// It handles the complete order flow: building the transaction, signing it, and submitting it.
// The operation wallet must be loaded before calling this method.
//
// Parameters:
//   - data: Order details including symbol, side, type, quantity, and optional price
//
// Returns:
//   - *SubmitPlaceOrderTransactionResponse: Order details and transaction info
//   - error: nil on success, error on failure
func (d *DeltaDeFi) PostOrder(data *BuildPlaceOrderTransactionRequest) (*SubmitPlaceOrderTransactionResponse, error) {
	if d.OperationWallet == nil {
		return nil, fmt.Errorf("operation wallet is not loaded")
	}

	buildRes, err := d.Order.BuildPlaceOrderTransaction(data)
	if err != nil {
		return nil, err
	}

	fmt.Println("Built order, tx hex:", buildRes.TxHex)
	signedTx, err := d.OperationWallet.Signer().SignTransaction(buildRes.TxHex)
	if err != nil {
		return nil, err
	}

	submitRes, err := d.Order.SubmitPlaceOrderTransaction(&SubmitPlaceOrderTransactionRequest{
		OrderID:  buildRes.OrderID,
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}

// CancelOrder is a high-level method for canceling an existing order.
// It handles the complete cancellation flow: building the transaction, signing it, and submitting it.
// The operation wallet must be loaded before calling this method.
//
// Parameters:
//   - orderId: The ID of the order to cancel
//
// Returns:
//   - *SubmitCancelOrderTransactionResponse: Transaction hash of the cancellation
//   - error: nil on success, error on failure
func (d *DeltaDeFi) CancelOrder(orderId string) (*SubmitCancelOrderTransactionResponse, error) {
	if d.OperationWallet == nil {
		return nil, fmt.Errorf("operation wallet is not loaded")
	}

	buildRes, err := d.Order.BuildCancelOrderTransaction(orderId)
	if err != nil {
		return nil, err
	}

	signedTx, err := d.OperationWallet.Signer().SignTransaction(buildRes.TxHex)
	if err != nil {
		return nil, err
	}

	submitRes, err := d.Order.SubmitCancelOrderTransaction(&SubmitCancelOrderTransactionRequest{
		SignedTx: signedTx,
	})
	if err != nil {
		return nil, err
	}
	return submitRes, nil
}

// CancelAllOrders is a high-level method for canceling all existing orders.
// It handles the complete cancellation flow: building the transaction, signing it, and submitting it.
// The operation wallet must be loaded before calling this method.
//
// Returns:
//   - *SubmitCancelAllOrdersTransactionResponse: Details of all canceled orders
//   - error: nil on success, error on failure
func (d *DeltaDeFi) CancelAllOrders() (*SubmitCancelAllOrdersTransactionResponse, error) {
	if d.OperationWallet == nil {
		return nil, fmt.Errorf("operation wallet is not loaded")
	}

	buildRes, err := d.Order.BuildCancelAllOrdersTransaction()
	if err != nil {
		return nil, err
	}

	signedTxs := make([]string, 0, len(buildRes.TxHexes))
	for _, txHex := range buildRes.TxHexes {
		signedTx, err := d.OperationWallet.Signer().SignTransaction(txHex)
		if err != nil {
			return nil, err
		}
		signedTxs = append(signedTxs, signedTx)
	}

	submitRes, err := d.Order.SubmitCancelAllOrdersTransaction(&SubmitCancelAllOrdersTransactionRequest{
		SignedTxs: signedTxs,
	})

	if err != nil {
		return nil, err
	}
	return submitRes, nil
}
