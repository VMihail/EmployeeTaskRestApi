create table if not exists Employees (
    id bigserial primary key not null,
    name varchar not null,
    email varchar not null unique
);

create table if not exists Task (
    id bigserial primary key not null,
    name varchar not null unique,
    description varchar not null
);

create table if not exists Employees_Task (
    id bigserial primary key not null,
    employeeId bigint not null, foreign key (employeeId) references Employees(id),
    taskId bigint not null, foreign key (taskId) references Task(id)
);