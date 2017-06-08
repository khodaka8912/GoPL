flag.CommandLineへvalueとしては*celsiusFlagを渡しており、celsiusFlagはCelsiusが埋め込まれているため、
デフォルト値の表示として(*celsiusFlag).String()->Celsius.String()が呼ばれるから。
（type Value interfaceはString()とSet()を要求している）