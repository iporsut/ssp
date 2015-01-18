ssp is Small Server Pipe

เอาไว้ทำให้เรียก http request เข้ามาแล้วส่งไปให้ command รัันแล้วเอา output ออกมา

ตัวอย่างเช่น ผมสร้างโปรแกรม add ไว้ดังนี้

```go
package main

import "fmt"

func addNumber(number1, number2 int) (result int) {
	return number1 + number2
}
func main() {
	var number1 int
	var number2 int
	fmt.Scanf("%d %d", &number1, &number2)
	var result = addNumber(number1, number2)
	fmt.Println(result)
}
```

ใช้ ssp สำหรับโปรแกรม add ดังนี้

ssp --cmd="add" --path="/add"

ผมสามารถยิง request มาได้ที่

http://localhost:8080/add?in=10 20

โดยใช้ query string "in" ในการส่ง input text เข้าไป
หรือถ้าเป็น POST จะใช้ request body ก็ได้

ถ้าต้องการเปลี่ยน port หรือ ip ของ host ให้ใช้ option --port และ --host

อีกตัวอย่างใช้กับคำสั่ง  sort ของ  shell

ssp --cmd="sort" --path="/sort"

ลองใช้ curl ยิง request ไปดังนี้
```shel
curl -X POST --data-binary @- http://127.0.0.1:8080/sort <<EOF
Weerasak
Kanokon
EOF
```

ได้ผลลัพธ์ออกมาดังนี้
```shell
Kanokon
Weerasak
```
