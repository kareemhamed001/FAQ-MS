package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	db "github.com/kareemhamed001/faq/internal/DB"
	"github.com/kareemhamed001/faq/internal/config"
	"github.com/kareemhamed001/faq/internal/seeders"
)

func main() {

	cfg := config.NewConfig()

	database, err := db.InitializeDB(cfg.DBDriver, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	seeder := seeders.NewSeeder(database)

	if len(os.Args) < 2 {
		showUsage()
		os.Exit(0)
	}

	command := os.Args[1]

	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	name := fs.String("name", "Admin", "Admin user name")
	email := fs.String("email", "admin@example.com", "Admin user email")
	password := fs.String("password", "admin123", "Admin user password")
	seedDefault := fs.Bool("default", false, "Seed default admin user (admin@example.com with password admin123)")

	fs.Parse(os.Args[2:])

	if command == "seed:admin" {
		fmt.Println("Seeding admin user...")
		if *seedDefault {
			if _, err := seeder.SeedDefaultAdminUser(); err != nil {
				log.Fatal("Error seeding admin user:", err)
			}
		} else {
			if _, err := seeder.SeedAdminUser(*name, *email, *password); err != nil {
				log.Fatal("Error seeding admin user:", err)
			}
		}
		fmt.Println("Admin user seeded successfully!")
		os.Exit(0)
	}

	showUsage()
}

func showUsage() {
	fmt.Println("\nCommands:")
	fmt.Println("  seed:admin    Seed admin user to database")
	fmt.Println("\nOptions:")
	fmt.Println("  -name       Admin user name (default: Admin)")
	fmt.Println("  -email      Admin user email (default: admin@example.com)")
	fmt.Println("  -password   Admin user password (default: admin123)")
	fmt.Println("  -default    Seed default admin user")
	fmt.Println("\nExamples:")
	fmt.Println("  go run cmd/seeder/main.go seed:admin -default")
	fmt.Println("  go run cmd/seeder/main.go seed:admin -name \"John Doe\" -email \"john@example.com\" -password \"secret123\"")
}
