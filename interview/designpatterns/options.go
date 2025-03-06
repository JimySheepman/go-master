package designpatterns

import "fmt"

// Config, yapılandırma seçeneklerini tutan yapı
type Config struct {
	Timeout    int
	MaxRetries int
	Verbose    bool
}

// Option, yapılandırma seçenekleri için fonksiyonel seçenek tipi
type Option func(*Config)

// WithTimeout, Timeout seçeneğini ayarlamak için bir seçenek fonksiyonu
func WithTimeout(timeout int) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithMaxRetries, MaxRetries seçeneğini ayarlamak için bir seçenek fonksiyonu
func WithMaxRetries(maxRetries int) Option {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithVerbose, Verbose seçeneğini ayarlamak için bir seçenek fonksiyonu
func WithVerbose(verbose bool) Option {
	return func(c *Config) {
		c.Verbose = verbose
	}
}

// NewConfig, Config yapısını oluşturan ve seçeneklerle yapılandıran fonksiyon
func NewConfig(opts ...Option) *Config {
	config := &Config{
		Timeout:    30,
		MaxRetries: 3,
		Verbose:    false,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}

func main() {
	// Timeout 50, MaxRetries 5 ve Verbose true olan bir Config oluşturuyoruz
	config := NewConfig(
		WithTimeout(50),
		WithMaxRetries(5),
		WithVerbose(true),
	)

	fmt.Println("Timeout:", config.Timeout)
	fmt.Println("Max Retries:", config.MaxRetries)
	fmt.Println("Verbose:", config.Verbose)
}
