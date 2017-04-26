package api

import "testing"

func TestDefaultConfiguration(t *testing.T) {
	c := DefaultConfiguration()

	if c.Key != "" {
		t.Errorf("Should have a empty Key")
	}
	if c.URL != DefaultURL {
		t.Errorf("Should have the default URL")
	}

	if c.Timeout != DefaultTimeout {
		t.Errorf("Should have the default Timeout")
	}

}

func TestDefaultConfigurationWithANewKey(t *testing.T) {
	key := "Key"
	c := DefaultConfiguration().WithKey(key)

	if c.Key != key {
		t.Errorf("Should have the same Key")
	}
	if c.URL != DefaultURL {
		t.Errorf("Should have the default URL")
	}

	if c.Timeout != DefaultTimeout {
		t.Errorf("Should have the default Timeout")
	}

}
