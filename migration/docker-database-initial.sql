create table category(
    id serial primary key,
    category_name varchar    
);


create table fuel(
    id serial primary key,
    fuel_name varchar
);

create table gas_station(
    id serial primary key,
    name varchar,
    category_id int DEFAULT NULL,
    FOREIGN KEY (category_id) REFERENCES category (id)
);



CREATE TABLE fuel_value(
    id serial primary key,
    fuel_id int DEFAULT NULL,
    gas_station_id int DEFAULT NULL,
    price float DEFAULT NULL,    
    FOREIGN KEY (fuel_id) REFERENCES fuel (id),
    FOREIGN KEY (gas_station_id) REFERENCES gas_station (id)
);

INSERT INTO category(category_name) VALUES
('Standard'),
('Deluxe'),
('Premium');

INSERT INTO fuel(fuel_name) VALUES
('Etanol'),
('Gasolina Comum'),
('Gasolina Adtivada'),
('Diesel');

INSERT INTO gas_station(name,category_id) VALUES
('Posto 1',1),
('Posto 2',2),
('Posto 3',3);


INSERT INTO fuel_value(fuel_id,gas_station_id,price) VALUES
(1,1,5.8),
(3,1,7.2),
(2,1,6.8),
(1,2,6.75);
