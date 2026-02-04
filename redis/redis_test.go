package redis

import (
	"os"
	"testing"
	"time"
)

func getTestClient() *RedisClient {
	addr := os.Getenv("TEST_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	password := os.Getenv("TEST_REDIS_PASSWORD")

	config := &Config{
		Addr:     addr,
		Password: password,
		DB:       0,
		PoolSize: 5,
	}
	return NewClientWithConfig(config)
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("Expected non-nil client")
	}
	defer client.Close()
}

func TestPing(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	pong, err := client.Ping()
	if err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
	if pong != "PONG" {
		t.Errorf("Expected PONG, got %s", pong)
	}
}

func TestSetGet(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	key := "test_key"
	value := "test_value"

	err := client.Set(key, value)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	retrieved, err := client.Get(key)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if retrieved != value {
		t.Errorf("Expected %s, got %s", value, retrieved)
	}

	// Cleanup
	client.Del(key)
}

func TestExists(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	key := "test_exists_key"

	exists, err := client.Exists(key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("Expected key to not exist")
	}

	client.Set(key, "value")

	exists, err = client.Exists(key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Error("Expected key to exist")
	}

	// Cleanup
	client.Del(key)
}

func TestTTL(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	key := "test_ttl_key"

	// Set key without TTL
	client.Set(key, "value")

	ttl, err := client.TTL(key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl != -1 {
		t.Errorf("Expected TTL -1, got %v", ttl)
	}

	// Set TTL
	client.Expire(key, 10)

	ttl, err = client.TTL(key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl <= 0 || ttl > 10*time.Second {
		t.Errorf("Expected TTL between 0 and 10s, got %v", ttl)
	}

	// Cleanup
	client.Del(key)
}

func TestExpire(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	key := "test_expire_key"

	client.Set(key, "value")

	err := client.Expire(key, 1) // 1 second
	if err != nil {
		t.Fatalf("Expire failed: %v", err)
	}

	time.Sleep(1100 * time.Millisecond) // Wait for expiration

	exists, err := client.Exists(key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("Expected key to be expired")
	}
}

func TestKeys(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	prefix := "test_keys_"

	// Set some keys
	client.Set(prefix+"1", "value1")
	client.Set(prefix+"2", "value2")
	client.Set("other_key", "value3")

	keys, err := client.Keys(prefix + "*")
	if err != nil {
		t.Fatalf("Keys failed: %v", err)
	}

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	// Cleanup
	client.Del(prefix + "1")
	client.Del(prefix + "2")
	client.Del("other_key")
}

func TestCountKeys(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	prefix := "test_count_"

	// Set some keys
	client.Set(prefix+"1", "value1")
	client.Set(prefix+"2", "value2")
	client.Set("other_key", "value3")

	count, err := client.CountKeys(prefix + "*")
	if err != nil {
		t.Fatalf("CountKeys failed: %v", err)
	}

	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}

	// Cleanup
	client.Del(prefix + "1")
	client.Del(prefix + "2")
	client.Del("other_key")
}

func TestScan(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	prefix := "test_scan_"

	// Set some keys
	for i := 0; i < 10; i++ {
		client.Set(prefix+string(rune('0'+i)), "value"+string(rune('0'+i)))
	}

	cursor := uint64(0)
	var allKeys []string

	for {
		keys, newCursor, err := client.Scan(cursor, prefix+"*", 5)
		if err != nil {
			t.Fatalf("Scan failed: %v", err)
		}
		allKeys = append(allKeys, keys...)
		cursor = newCursor
		if cursor == 0 {
			break
		}
	}

	if len(allKeys) != 10 {
		t.Errorf("Expected 10 keys, got %d", len(allKeys))
	}

	// Cleanup
	for i := 0; i < 10; i++ {
		client.Del(prefix + string(rune('0'+i)))
	}
}

func TestDel(t *testing.T) {
	client := getTestClient()
	defer client.Close()

	key := "test_del_key"

	client.Set(key, "value")

	err := client.Del(key)
	if err != nil {
		t.Fatalf("Del failed: %v", err)
	}

	exists, err := client.Exists(key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("Expected key to be deleted")
	}
}
