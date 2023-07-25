CREATE USER 'npiaorchestrator'@'%' IDENTIFIED BY 'youdonthavetoknow';

GRANT ALL PRIVILEGES ON *.* TO 'npiaorchestrator'@'%';

CREATE DATABASE orchestrator;

USE orchestrator;

CREATE TABLE orchestrator_record (email VARCHAR(100), config TEXT, osid VARCHAR(50), request_key VARCHAR(100));


SET @config_amld_enc = LOAD_FILE("/var/lib/mysql-files/config.amld");


INSERT INTO orchestrator_record(email, config, osid, request_key) VALUES("seantywork@gmail.com",@config_amld_enc, "N","N");




COMMIT;