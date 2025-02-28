package instaclustr

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAPIClientDeleteKafkaAcl(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/acls", `{"id":"should-be-uuid"}`, 200)
	err := client.DeleteKafkaAcl("should-be-uuid", nil)
	if err != nil {
		t.Fatalf("Failed to delete kafka ACL: %s", err)
	}
}

func TestAPIClientCreateKafkaAcl(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/acls", `{"id":"should-be-uuid"}`, 200)
	err2 := client.CreateKafkaAcl("should-be-uuid", nil)
	if err2 != nil {
		t.Fatalf("Failed to create kafka ACL: %s", err2)
	}
}

func TestAPIClientReadKafkaAcls(t *testing.T) {
	filename := "data/valid_kafka_acls.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/acls/searches", jsonStr, 200)
	acls, err2 := client.ReadKafkaAcls("should-be-uuid", nil)
	if err2 != nil {
		t.Fatalf("Failed to list Kafka ACL: %s", err2)
	}
	if acls[0].Principal != "User:test1" || acls[1].Principal != "User:test2" {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientCreateKafkaTopic(t *testing.T) {
	filename := "data/valid_kafka_topic_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/kafka/topics", `{"id":"should-be-uuid"}`, 201)
	err2 := client.CreateKafkaTopic("should-be-uuid", jsonStr)
	if err2 != nil {
		t.Fatalf("Failed to create kafka topic: %s", err2)
	}
}

func TestAPIDeleteKafkaTopic(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/topics/test", `{"id":"should-be-uuid"}`, 200)
	err := client.DeleteKafkaTopic("should-be-uuid", "test")
	if err != nil {
		t.Fatalf("Failed to delete kafka topic: %s", err)
	}
}

func TestAPIClientReadKafkaTopicConfig(t *testing.T) {
	filename := "data/valid_kafka_topic_config.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/topics/test/config", jsonStr, 200)
	values, err2 := client.ReadKafkaTopicConfig("should-be-uuid", "test")
	if err2 != nil {
		t.Fatalf("Failed to read Kafka topic config: %s", err2)
	}
	if (*values).Config.CompressionType != "producer" || *(*values).Config.MessageDownconversionEnable != true ||
		(*values).Config.MinInsyncReplicas != 2 {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientCreateKafkaTopicList(t *testing.T) {
	filename := "data/valid_kafka_topic_list.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/topics", jsonStr, 200)
	topicList, err2 := client.ReadKafkaTopicList("should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to create kafka topic: %s", err2)
	}
	if topicList.Topics[0] != "test1" || topicList.Topics[1] != "test2" {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientReadKafkaTopic(t *testing.T) {
	filename := "data/valid_kafka_topic_read.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/topics/test", jsonStr, 200)
	values, err2 := client.ReadKafkaTopic("should-be-uuid", "test")
	if err2 != nil {
		t.Fatalf("Failed to read Kafka topic config: %s", err2)
	}
	if (*values).Topic != "test" || (*values).ReplicationFactor != 3 ||
		(*values).Partitions != 3 {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientCreateKafkaUser(t *testing.T) {
	filename := "data/valid_kafka_user.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/kafka/users", `{"id":"should-be-uuid"}`, 201)
	err2 := client.CreateKafkaUser("should-be-uuid", jsonStr)
	if err2 != nil {
		t.Fatalf("Failed to create kafka user: %s", err2)
	}
}

func TestAPIClientCreateKafkaUserList(t *testing.T) {
	filename := "data/valid_kafka_user_list.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/users", jsonStr, 200)
	topicList, err2 := client.ReadKafkaUserList("should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to create kafka topic: %s", err2)
	}
	if topicList[0] != "test1" || topicList[1] != "test2" || topicList[2] != "test3" {
		t.Fatalf("Values do not match.")
	}
}
