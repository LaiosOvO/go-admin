
#### viper
1 定义viper读取配置文件
2 将viper的内容反序列化到一个对象里面
```go
    
    v := viper.New()
	v.SetConfigFile("配置文件的名称以及路径")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}
```

#### integrate gin



#### integrate gorm




#### integrate db-list
- 读取配置文件信息
- 整合多个数据库的源放进 dbMap里面

用gorm打开一个数据库实例
```go
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
```

##### psql

1. 数据源配置类 config/gorm_pgsql initialize/gorm_pgsql
   链接的psql返回db的实例
```go
func GormPgSql() *gorm.DB {
	p := global.GVA_CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
```


#### 整合mongodb
- 配置类
- 直接通过 qmgo 来获取一个数据库的实例

```go
func (m *mongo) Initialization() error {
	var opts []options.ClientOptions
	if global.GVA_CONFIG.Mongo.IsZap {
		opts = internal.Mongo.GetClientOptions()
	}
	ctx := context.Background()
	client, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:              global.GVA_CONFIG.Mongo.Uri(),
		Coll:             global.GVA_CONFIG.Mongo.Coll,
		Database:         global.GVA_CONFIG.Mongo.Database,
		MinPoolSize:      &global.GVA_CONFIG.Mongo.MinPoolSize,
		MaxPoolSize:      &global.GVA_CONFIG.Mongo.MaxPoolSize,
		SocketTimeoutMS:  &global.GVA_CONFIG.Mongo.SocketTimeoutMs,
		ConnectTimeoutMS: &global.GVA_CONFIG.Mongo.ConnectTimeoutMs,
		Auth: &qmgo.Credential{
			Username:   global.GVA_CONFIG.Mongo.Username,
			Password:   global.GVA_CONFIG.Mongo.Password,
			AuthSource: global.GVA_CONFIG.Mongo.AuthSource,
		},
	}, opts...)
	if err != nil {
		return errors.Wrap(err, "链接mongodb数据库失败!")
	}
	global.GVA_MONGO = client
	err = m.Indexes(ctx)
	if err != nil {
		return err
	}
	return nil
}
```


