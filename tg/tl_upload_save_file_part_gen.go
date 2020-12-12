// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// UploadSaveFilePartRequest represents TL type `upload.saveFilePart#b304a621`.
// Saves a part of file for futher sending to one of the methods.
//
// See https://core.telegram.org/method/upload.saveFilePart for reference.
type UploadSaveFilePartRequest struct {
	// Random file identifier created by the client
	FileID int64
	// Numerical order of a part
	FilePart int
	// Binary data, contend of a part
	Bytes []byte
}

// UploadSaveFilePartRequestTypeID is TL type id of UploadSaveFilePartRequest.
const UploadSaveFilePartRequestTypeID = 0xb304a621

// Encode implements bin.Encoder.
func (s *UploadSaveFilePartRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveFilePart#b304a621 as nil")
	}
	b.PutID(UploadSaveFilePartRequestTypeID)
	b.PutLong(s.FileID)
	b.PutInt(s.FilePart)
	b.PutBytes(s.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (s *UploadSaveFilePartRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveFilePart#b304a621 to nil")
	}
	if err := b.ConsumeID(UploadSaveFilePartRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field file_id: %w", err)
		}
		s.FileID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field file_part: %w", err)
		}
		s.FilePart = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveFilePart#b304a621: field bytes: %w", err)
		}
		s.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadSaveFilePartRequest.
var (
	_ bin.Encoder = &UploadSaveFilePartRequest{}
	_ bin.Decoder = &UploadSaveFilePartRequest{}
)

// UploadSaveFilePart invokes method upload.saveFilePart#b304a621 returning error if any.
// Saves a part of file for futher sending to one of the methods.
//
// See https://core.telegram.org/method/upload.saveFilePart for reference.
func (c *Client) UploadSaveFilePart(ctx context.Context, request *UploadSaveFilePartRequest) (bool, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
