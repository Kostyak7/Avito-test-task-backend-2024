-- Создание роли и базы данных
CREATE ROLE myuser WITH LOGIN PASSWORD 'rew12345';
CREATE DATABASE avitotestdbname;
GRANT ALL PRIVILEGES ON DATABASE avitotestdbname TO myuser;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

CREATE TABLE tender (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    service_type VARCHAR(50) CHECK (service_type IN ('Construction', 'Delivery', 'Manufacture')),
    status VARCHAR(50) CHECK (status IN ('Created', 'Published', 'Closed')),
    organization_id UUID NOT NULL,
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    creator_username VARCHAR(100) NOT NULL
);

CREATE TABLE bids (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('Created', 'Published', 'Canceled')),
    tender_id UUID NOT NULL,
    author_type VARCHAR(50) CHECK (author_type IN ('Organization', 'User')),
    author_id UUID NOT NULL,
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bid_reviews (
    id UUID PRIMARY KEY,
    description VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    bid_id UUID NOT NULL REFERENCES bids(id)
);