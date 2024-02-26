package config

type AppConfig struct {
	ServiceName string             `mapstructure:"serviceName"`
	Development bool               `mapstructure:"development"`
	Logger      *LoggerConfig      `mapstructure:"LOGGER"`
	Http        *HttpConfig        `mapstructure:"http"`
	GRPC        *GrpcConfig        `mapstructure:"grpc"`
	Database    *DatabaseConfig    `mapstructure:"DATABASE"`
	Jaeger      *JaegerConfig      `mapstructure:"jaeger"`
	Kafka       *KafkaConfig       `mapstructure:"KAFKA"`
	Redis       *RedisConfig       `mapstructure:"REDIS"`
	Healthcheck *HealthcheckConfig `mapstructure:"healthcheck"`
	Metrics     *MetricsConfig     `mapstructure:"metrics"`
	Scheduler   *SchedulerConfig   `mapstructure:"SCHEDULER"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"LOGGER_LEVEL"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

type HttpConfig struct {
	Port            string `mapstructure:"port"`
	Development     bool   `mapstructure:"development"`
	ShutdownTimeout int    `mapstructure:"shutdownTimeout"`

	Resources    []string            `mapstructure:"resources"`
	RateLimiting *RateLimitingConfig `mapstructure:"rateLimiting"`
}

type RateLimitingConfig struct {
	RateFormat string `mapstructure:"rateFormat"`
}

type GrpcConfig struct {
	Port              string `mapstructure:"port"`
	Development       bool   `mapstructure:"development"`
	MaxConnectionIdle int    `mapstructure:"maxConnectionIdle"`
	Timeout           int    `mapstructure:"timeout"`
	MaxConnectionAge  int    `mapstructure:"maxConnectionAge"`
	Time              int    `mapstructure:"time"`
}
type DatabaseConfig struct {
	ReadDbCfg  *ReadDbConfig  `mapstructure:"READDB"`
	WriteDbCfg *WriteDbConfig `mapstructure:"WRITEDB"`
}

type ReadDbConfig struct {
	DbType            string `mapstructure:"dbType"`
	ConnectionString  string `mapstructure:"CONNECTION_STRING"`
	MigrationFilePath string `mapstructure:"migrationFilePath"`
}

type WriteDbConfig struct {
	DbType            string `mapstructure:"dbType"`
	ConnectionString  string `mapstructure:"CONNECTION_STRING"`
	MigrationFilePath string `mapstructure:"migrationFilePath"`
}

type JaegerConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	HostPort    string `mapstructure:"hostPort"`
	Enable      bool   `mapstructure:"enable"`
	LogSpans    bool   `mapstructure:"logSpans"`
}

type KafkaConfig struct {
	Config    *KafkaConfigDetail `mapstructure:"CONFIG"`
	Consumers *ConsumerConfig    `mapstructure:"consumers"`
	Producers *ProducerConfig    `mapstructure:"producers"`
}

type KafkaConfigDetail struct {
	Brokers  []string `mapstructure:"BROKERS"`
	Username string   `mapstructure:"USERNAME"`
	Password string   `mapstructure:"PASSWORD"`
}

type ConsumerConfig struct {
	GemCreateConsumer *GemCreateConsumerConfig `mapstructure:"gemCreateConsumer"`
}

type ProducerConfig struct {
	GemCreateProducer *GemCreateProducerConfig `mapstructure:"gemCreateProducer"`
}

type GemCreateConsumerConfig struct {
	GroupId   string `mapstructure:"groupId"`
	TopicName string `mapstructure:"topicName"`
	NumWorker int    `mapstructure:"NumWorker"`
}

type GemCreateProducerConfig struct {
	TopicName         string `mapstructure:"topicName"`
	InitTopic         bool   `mapstructure:"initTopic"`
	NumPartitions     int    `mapstructure:"numPartitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}

type RedisConfig struct {
	Addrs    []string `mapstructure:"ADDRESS"`
	Password string   `mapstructure:"PASSWORD"`
	DB       int      `mapstructure:"db"`
	PoolSize int      `mapstructure:"poolSize"`
}

type HealthcheckConfig struct {
	Interval           int    `mapstructure:"interval"`
	Port               string `mapstructure:"port"`
	GoroutineThreshold int    `mapstructure:"goroutineThreshold"`
}

type MetricsConfig struct {
	PrometheusPath string `mapstructure:"prometheusPath"`
	PrometheusPort string `mapstructure:"prometheusPort"`
}

type SchedulerConfig struct {
	CronExpression string `mapstructure:"CRON_EXPRESSION"`
}
