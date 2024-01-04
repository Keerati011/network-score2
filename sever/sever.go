package main

//นำเข้าเเพคเกจที่จะใช้
import (
	"fmt" // พิมพ์ข้อความออกทางหน้าจอ
	"net" // เชื่อมต่อเครือข่าย
)

// 4.สร้างฟังชันจัดการเชื่อต่อ
// รับพารามิเตอร์เป็นตัวแปร conn ชนิด net.Conn
func handleConnection(conn net.Conn) {
	defer conn.Close() //ปิดการเชื่อมต่อ เมื่อ handleConnection จบการทำงาน
	// สร้างตัวแปร buffer เก็บข้อมูล
	buffer := make([]byte, 1024) //สร้าง buffer ชนิด byte ความจุ 1024
	//รอรับข้อมูลจากตัวแปร และอ่าน buffer โดย ฟังชัน conn.Read
	//ส่งจำนวน
	buffer1 := make([]byte, 1024)
	username := string(buffer)
	password := string(buffer1)
	for {
		n, err := conn.Read(buffer)
		m, err := conn.Read(buffer1)
		// ตรวจ err
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		//  แสดงข้อความออกทางหน้าจอ
		fmt.Println("username:", string(buffer[:n]))
		fmt.Println("password:", string(buffer1[:m]))

		if username == "std1" {
			if password == "p@ssw0rd" {
				conn.Write([]byte("[Login is successful]\n"))
				conn.Write([]byte("sever response: hello\n"))
			}
		} else {
			conn.Write([]byte("[Login is unsuccessful]\n"))
			conn.Write([]byte("sever response: Invalid credentials\n"))
		}

	}
}

// ฟังชัน main เป็นฟังชันหลักของโปรแกรม
func main() {
	//1.สร้างตัวแปร เพื่อรอรับการเชื่อมต่อเน็ต
	//port 5000
	listener, err := net.Listen("tcp", ":5000")
	//2.เชค error ไม่เท่ากับ ค่าว่าง ถ้าใช่แสดงข้อความ err
	if err != nil {
		fmt.Println(err)
		// return จบการทำงงาน
		return
	}
	//ปิดการเชื่อมต่อ เมื่อ  main จบการทำงาน
	defer listener.Close()
	//แสดงข้อความว่ารอรับอยุ่
	fmt.Println("Server is listening on port 5000")

	//3.รอรับการเชื่อมต่อ
	for {
		// รับการเชื่อมต่อ
		conn, err := listener.Accept() // conn ตัวแปรที่บอกว่าเชื่อมต่ออยุ่
		//ตรวจสอบ err
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		fmt.Println("New connection established")
		// ส่งตัวแปร conn ไป ฟังชัน handleConnection
		go handleConnection(conn)
	}
}
