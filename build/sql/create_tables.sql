CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,                     
    phone_number TEXT,                             
    description TEXT,                              
    user_pic TEXT,                                 
    password_hash BYTEA NOT NULL                   
);

CREATE TABLE IF NOT EXISTS restaurants (
    id UUID PRIMARY KEY,  
    name TEXT NOT NULL,                            
    description TEXT,                              
    type TEXT,                                     
    rating FLOAT                                   
);