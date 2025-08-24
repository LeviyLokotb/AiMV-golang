package data

const TZ = `
+---------------+-------+-------+
|team:          |   B1  |   B2  |
+---------------+-------+-------+
|per hour:      |       |       |	
|   I1:         |   4   |   1   |
|   I2:	        |   2   |   3   |
+---------------+-------+-------+
|time (hours):  |  9.5  |   4   |
+---------------+-------+-------+
|price (1 item):|       |       |
|   I1:	        |   9   |  15   |
|   I2:	        |  20   |  30   |
+---------------+-------+-------+

Найти оптимальный обхём выпуска изделий, обеспечивающий 
минимальные затраты на выполнение заказа.
`

type Team struct {
	I1ph     float64
	I2ph     float64
	Worktime float64
	I1price  float64
	I2price  float64
}

var B1 Team = Team{
	I1ph:     4,
	I2ph:     2,
	Worktime: 9.5,
	I1price:  9.0,
	I2price:  20.0,
}

var B2 Team = Team{
	I1ph:     1,
	I2ph:     3,
	Worktime: 4.0,
	I1price:  15.0,
	I2price:  30.0,
}

var TOTAL_I1 = 32.0
var TOTAL_I2 = 4.0
