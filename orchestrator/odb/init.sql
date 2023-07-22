CREATE USER 'npiaorchestrator'@'%' IDENTIFIED BY 'youdonthavetoknow';

GRANT ALL PRIVILEGES ON *.* TO 'npiaorchestrator'@'%';

CREATE DATABASE orchestrator;

USE orchestrator;

CREATE TABLE orchestrator_record (email VARCHAR(100), config TEXT, signed TINYINT);


SET @config_amld_enc = LOAD_FILE("/var/lib/mysql-files/config.amld");


INSERT INTO orchestrator_record(email, config) VALUES("seantywork@gmail.com",@config_amld_enc);




COMMIT;