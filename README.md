# Go com MySQL

## Rodando MySQL em Docker

Comando para iniciar um Docker MySQL

```sh
docker run \
  --rm \
  --name mysql \
  -d \
  -e MYSQL_ROOT_PASSWORD=Prud@123 \
  -e MYSQL_DATABASE=golang \
  -p 3306:3306 \
  mysql:8.0.27
  ```

## Comando para abrir um shell no Docker MySQL e executar comandos SQL

```sh
docker exec -it mysql /bin/bash
mysql -u root -p
```

## Comando para trocar a senha do root

```sql
ALTER USER 'root'@'localhost' IDENTIFIED BY 'S3nh4C0mpl3x4$';
FLUSH PRIVILEGES;
```

## Criação de um usuário de serviço

```sql
CREATE USER 'golang'@'%' IDENTIFIED BY 'golang';
ALTER USER 'golang'@'%' IDENTIFIED BY 'S3nh4C0mpl3x4$';
GRANT all privileges on golang.* to 'golang'@'%';
FLUSH PRIVILEGES;
```

## Criação de uma tabela de exemplo, usando o usuário de serviço

```sh
docker exec -it mysql /bin/bash
mysql -u golang -p
```

Criação de uma tabela pra teste

```sql
USE golang;

CREATE TABLE produto (
  cd_produto INT AUTO_INCREMENT PRIMARY KEY,
  ds_produto VARCHAR(100) NOT NULL
);
```

Incluão de dados...

```sql
USE golang;

INSERT INTO produto (ds_produto) values ('iPhone 13 Pro');
INSERT INTO produto (ds_produto) values ('iPhone 13 Pro Max');
INSERT INTO produto (ds_produto) values ('AirPod Pro');
INSERT INTO produto (ds_produto) values ('AirPod Max');
```
