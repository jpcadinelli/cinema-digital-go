package conection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	DB, err = gorm.Open(postgres.Open(getStringConection()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	log.Println("Connected to database")
}

func getStringConection() string {
	dns := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	return dns
}

func GetConnection() (*gorm.DB, error) {
	DB, err = gorm.Open(postgres.Open(getStringConection()), &gorm.Config{})
	if err != nil {
		return DB, fmt.Errorf("Could not connect to database: %v", err)
	}
	return DB, nil
}

func RunMigrations() error {
	migrationsDir := "api/pkg/database/migrations"

	err = createMigrationHistoryTable()
	if err != nil {
		return fmt.Errorf("erro ao criar a tabela de histórico de migrações: %w", err)
	}

	var files []os.DirEntry

	files, err = os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("erro ao ler o diretório de migrações: %w", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		sqlFiles = append(sqlFiles, filepath.Join(migrationsDir, file.Name()))
	}

	for _, file := range sqlFiles {
		if isMigrationApplied(file) {
			log.Printf("Migração %s já foi aplicada. Pulando...", file)
			continue
		}

		if err = runSQLScript(file); err != nil {
			return fmt.Errorf("erro ao executar migração do arquivo %s: %w", file, err)
		}

		if err = registerMigration(file); err != nil {
			return fmt.Errorf("erro ao registrar migração %s: %w", file, err)
		}
	}

	return nil
}

func isMigrationApplied(migrationName string) bool {
	var count int64
	err = DB.Raw("SELECT COUNT(*) FROM migration_history WHERE migration_name = ?", migrationName).Scan(&count).Error
	if err != nil {
		log.Fatal("Erro ao verificar migração:", err)
	}
	return count > 0
}

func registerMigration(migrationName string) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err = tx.Exec("INSERT INTO migration_history (migration_name, applied_at) VALUES (?, ?)", migrationName, time.Now()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func createMigrationHistoryTable() error {
	return DB.Exec(`
		CREATE TABLE IF NOT EXISTS migration_history (
			id SERIAL PRIMARY KEY,
			migration_name TEXT NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`).Error
}

func runSQLScript(scriptPath string) error {
	var sqlBytes []byte

	sqlBytes, err = os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo SQL %s: %w", scriptPath, err)
	}

	sqlScript := string(sqlBytes)

	tx := DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("erro ao iniciar a transação: %w", tx.Error)
	}

	if err = tx.Exec(sqlScript).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao executar o script SQL %s: %w", scriptPath, err)
	}

	if err = tx.Commit().Error; err != nil {
		return fmt.Errorf("erro ao commitar a transação: %w", err)
	}

	log.Printf("Migração executada com sucesso: %s", scriptPath)
	return nil
}
