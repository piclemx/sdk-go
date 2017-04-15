package discovery

import "testing"

func TestDefaultConfiguration(t *testing.T) {
	c := DefaultConfiguration()

	if c.key != "" {
		t.Errorf("Should have a empty key")
	}
	if c.url != DefaultURL {
		t.Errorf("Should have the default URL")
	}

	if c.timeout != DefaultTimeout {
		t.Errorf("Should have the default timeout")
	}

}

func TestDefaultConfigurationWithANewKey(t *testing.T) {
	key := "key"
	c := DefaultConfiguration().WithKey(key)

	if c.key != key {
		t.Errorf("Should have the same key")
	}
	if c.url != DefaultURL {
		t.Errorf("Should have the default URL")
	}

	if c.timeout != DefaultTimeout {
		t.Errorf("Should have the default timeout")
	}

}
