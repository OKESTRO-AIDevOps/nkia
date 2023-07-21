CREATE USER 'npiaorchestrator'@'%' IDENTIFIED BY 'youdonthavetoknow';

GRANT ALL PRIVILEGES ON *.* TO 'npiaorchestrator'@'%';

CREATE DATABASE orchestrator;

USE orchestrator;

CREATE TABLE orchestrator_pair (email_entry VARCHAR(100), amalgamated_config TEXT);



SET @config_amld_enc = LOAD_FILE("/var/lib/mysql-files/config.amld");


INSERT INTO orchestrator_pair(email_entry, amalgamated_config) VALUES("seantywork@gmail.com",@config_amld_enc);




COMMIT;