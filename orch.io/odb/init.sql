CREATE USER 'npiaorchestrator'@'%' IDENTIFIED BY 'youdonthavetoknow';

GRANT ALL PRIVILEGES ON *.* TO 'npiaorchestrator'@'%';

CREATE DATABASE orchestrator;

USE orchestrator;

CREATE TABLE orchestrator_record (email VARCHAR(100), pubkey TEXT, osid VARCHAR(50), request_key VARCHAR(100), PRIMARY KEY(email));

CREATE TABLE orchestrator_cluster_record (email VARCHAR(100), cluster_id VARCHAR(100), config TEXT, config_status VARCHAR(10), PRIMARY KEY(email, cluster_id));

-- SET @config_amld_enc = LOAD_FILE("/var/lib/mysql-files/config.amld");


COMMIT;