package main

import (
	"fmt"
	"time"
)

// AttendanceManager interface untuk mengelola absensi karyawan
type AttendanceManager interface {
	RecordAttendance(employeeID int)
}

// Employee struct untuk merepresentasikan data karyawan
type Employee struct {
	ID        int
	Name      string
	Position  string
	CreatedAt time.Time
}

// EmployeeRepository struct untuk mengelola data karyawan
type EmployeeRepository struct {
	employees []Employee
}

// AddEmployee metode untuk menambahkan karyawan baru
func (er *EmployeeRepository) AddEmployee(employee Employee) {
	er.employees = append(er.employees, employee)
}

// RemoveEmployee metode untuk menghapus karyawan berdasarkan ID
func (er *EmployeeRepository) RemoveEmployee(employeeID int) {
	for i, employee := range er.employees {
		if employee.ID == employeeID {
			er.employees = append(er.employees[:i], er.employees[i+1:]...)
			break
		}
	}
}

// FindEmployeeByID metode untuk mencari karyawan berdasarkan ID
func (er *EmployeeRepository) FindEmployeeByID(employeeID int) *Employee {
	for _, employee := range er.employees {
		if employee.ID == employeeID {
			return &employee
		}
	}
	return nil
}

// AttendanceService struct untuk mengelola absensi karyawan
type AttendanceService struct {
	employeeRepository *EmployeeRepository
}

// ClockIn metode untuk melakukan absensi masuk
func (as *AttendanceService) ClockIn(employeeID int) {
	employee := as.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi masuk berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

// ClockOut metode untuk melakukan absensi keluar
func (as *AttendanceService) ClockOut(employeeID int) {
	employee := as.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi keluar berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

// EmptyAttendanceService struct untuk mengelola absensi karyawan dengan metode absensi kosong
type EmptyAttendanceService struct{}

// ClockIn metode untuk melakukan absensi masuk (implementasi kosong)
func (eas *EmptyAttendanceService) ClockIn(employeeID int) {
	fmt.Println("Metode tidak tersedia")
}

// ClockOut metode untuk melakukan absensi keluar (implementasi kosong)
func (eas *EmptyAttendanceService) ClockOut(employeeID int) {
	fmt.Println("Metode tidak tersedia")
}

func main() {
	employeeRepo := &EmployeeRepository{}
	attendanceService := &AttendanceService{employeeRepository: employeeRepo}

	emptyAttendanceService := &EmptyAttendanceService{}

	employee1 := Employee{ID: 1, Name: "John Doe", Position: "Manager", CreatedAt: time.Now()}
	employee2 := Employee{ID: 2, Name: "Jane Smith", Position: "Staff", CreatedAt: time.Now()}

	employeeRepo.AddEmployee(employee1)
	employeeRepo.AddEmployee(employee2)

	attendanceService.ClockIn(1)
	attendanceService.ClockIn(2)

	attendanceService.ClockOut(1)
	attendanceService.ClockOut(2)

	// Penggunaan objek EmptyAttendanceService sebagai pengganti AttendanceService
	emptyAttendanceService.ClockIn(1)
	emptyAttendanceService.ClockOut(2)
}
