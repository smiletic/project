insert into role values(1,'admin');
insert into role values(2,'doctor');
insert into role values(3,'researcher');
insert into role values(4,'nurse');
insert into specialty values(1,'common medicine');
insert into specialty values(2,'psychologist');
insert into specialty values(3,'ortophedist');
insert into person(uid,name,surname,JMBG) values ('0d60a85e-0b90-4482-a14c-108aea2557aa', 'Peter', 'Parker', '1223123123123');
insert into person(uid,name,surname,JMBG) values ('39240e9f-ae09-4e95-9fd0-a712035c8ad7', 'Bruce', 'Banner', '2352352352352');
insert into person(uid,name,surname,JMBG) values ('9e4de779-d6a0-44bc-a531-20cdb97178d2', 'Thor', 'Odinson', '9112491249994');
insert into person(uid,name,surname,JMBG) values ('66a45c1b-19af-4ab5-8747-1b0e2d79339d', 'Natasha', 'Romanova', '255422211234');

insert into employee values ('0d60a85e-0b90-4482-a14c-108aea2557aa', '0d60a85e-0b90-4482-a14c-108aea2557aa', '1223123123123');
insert into employee values ('39240e9f-ae09-4e95-9fd0-a712035c8ad7', '39240e9f-ae09-4e95-9fd0-a712035c8ad7', '2352352352352');
insert into employee values ('9e4de779-d6a0-44bc-a531-20cdb97178d2', '9e4de779-d6a0-44bc-a531-20cdb97178d2', '9112491249994');
insert into employee values ('66a45c1b-19af-4ab5-8747-1b0e2d79339d', '66a45c1b-19af-4ab5-8747-1b0e2d79339d', '255422211234');

insert into system_user(uid,employee_uid,role,username,password) values ('0d60a85e-0b90-4482-a14c-108aea2557aa','0d60a85e-0b90-4482-a14c-108aea2557aa',1,'spiderman','9f05aa4202e4ce8d6a72511dc735cce9');
insert into system_user(uid,employee_uid,role,username,password) values ('39240e9f-ae09-4e95-9fd0-a712035c8ad7','39240e9f-ae09-4e95-9fd0-a712035c8ad7',2,'hulk','76254239879f7ed7d73979f1f7543a7e');
insert into system_user(uid,employee_uid,role,username,password) values ('9e4de779-d6a0-44bc-a531-20cdb97178d2','9e4de779-d6a0-44bc-a531-20cdb97178d2',3,'thor','575e22bc356137a41abdef379b776dba');
insert into system_user(uid,employee_uid,role,username,password) values ('66a45c1b-19af-4ab5-8747-1b0e2d79339d','66a45c1b-19af-4ab5-8747-1b0e2d79339d',4,'widow','C9AD31E5740747285DAE5C168715D2DE');
