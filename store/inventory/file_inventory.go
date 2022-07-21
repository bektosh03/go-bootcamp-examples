package inventory

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"store/product"
	"strconv"
	"strings"
)

type FileInventory struct {
	db       *os.File
	products product.List
}

func NewFileInventory(name string) (*FileInventory, error) {
	f, err := os.OpenFile(name, os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	i := &FileInventory{db: f}
	if err = i.load(); err != nil {
		return nil, err
	}

	return i, nil
}

func (i *FileInventory) Close() error {
	return i.db.Close()
}



func (i *FileInventory) FindProduct(name string) (product.Product, bool) {
	for _, p := range i.products {
		if p.Name == name {
			return p, true
		}
	}

	return product.Product{}, false
}

func (i *FileInventory) AddProduct(p product.Product) error {
	i.products.Add(p)
	if err := i.snapshot(); err != nil {
		i.products.Remove(p.Name)
		return err
	}

	return nil
}

func (i *FileInventory) load() error {
	scanner := bufio.NewScanner(i.db)
	for scanner.Scan() {
		p, err := i.parseLine(scanner.Text())
		if err != nil {
			return fmt.Errorf("failed to parse line: %w", err)
		}

		i.products = append(i.products, p)
	}

	return nil
}

func (i *FileInventory) parseLine(line string) (product.Product, error) {
	lineItems := strings.Fields(line)
	if len(lineItems) < 4 {
		return product.Product{}, fmt.Errorf("invalid data format: not enough items: %d", len(lineItems))
	}
	quantity, err := strconv.ParseUint(lineItems[1], 10, 64)
	if err != nil {
		return product.Product{}, fmt.Errorf("invalid data at 1: expected uint64, got %T-%v", lineItems[1], lineItems[1])
	}
	price, err := strconv.ParseUint(lineItems[2], 10, 64)
	if err != nil {
		return product.Product{}, fmt.Errorf("invalid data at 2: expected uint64, got %T-%v", lineItems[2], lineItems[2])
	}
	originalPrice, err := strconv.ParseUint(lineItems[3], 10, 64)
	if err != nil {
		return product.Product{}, fmt.Errorf("invalid data at 3: expected uint64, got %T-%v", lineItems[3], lineItems[3])
	}

	return product.Product{
		Name:          lineItems[0],
		Quantity:      quantity,
		Price:         price,
		OriginalPrice: originalPrice,
	}, nil
}

func (i *FileInventory) snapshot() error {
	if err := os.Truncate(i.db.Name(), 0); err != nil {
		return fmt.Errorf("failed to truncate: %w", err)
	}
	if _, err := i.db.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek: %w", err)
	}

	var data strings.Builder
	for _, p := range i.products {
		if _, err := data.WriteString(fmt.Sprintf("%s %d %d %d\n", p.Name, p.Quantity, p.Price, p.OriginalPrice)); err != nil {
			return err
		}
	}

	_, err := i.db.WriteString(data.String())
	return err
}
