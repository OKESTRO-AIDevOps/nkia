CREATE USER 'npiaorchestrator'@'%' IDENTIFIED BY 'youdonthavetoknow';

GRANT ALL PRIVILEGES ON *.* TO 'npiaorchestrator'@'%';

CREATE DATABASE orchestrator;

USE orchestrator;


COMMIT;