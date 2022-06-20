package inventory

import (
	"bufio"
	"fmt"
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
