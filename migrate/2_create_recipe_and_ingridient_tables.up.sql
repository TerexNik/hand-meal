create table recipe
(
    id          int  not null generated always as identity unique primary key,
    title       text not null,
    description text,
    calories    int,
    time        timestamp,
    image_url   text,
    author_id   int references users (id)
);

create table ingredient
(
    id          int  not null generated always as identity unique primary key,
    name        text not null,
    description text,
    weight      int,
    units       varchar(40)
);

create table recipe_ingredient
(
    recipe_id  int references recipe (id),
    ingredient int references ingredient (id)
);