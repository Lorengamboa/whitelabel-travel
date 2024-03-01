-- migrations/000003_create_travel_packages_table.up
-- Add up migration script here

CREATE TABLE travel_packages (
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  client_id UUID NOT NULL,
  package_name VARCHAR(255) NOT NULL,
  duration INTEGER NOT NULL,
  itinerary TEXT NOT NULL,
  package_includes TEXT NOT NULL,
  package_excludes TEXT NOT NULL,
  recommended_gear TEXT,
  difficulty_level TEXT,
  price DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);