export interface ClientBody {
  url: string;
  title: string;
  primaryColor: string;
  secondaryColor: string;
  logo: string;
  name: string;
  email: string;
  address1: string;
  address2: string;
  city: string;
  country: string;
}

export interface TravelPackageBody {
  travelPackage: string;
  duration: number;
  itinerary: string;
  packageIncludes: string;
  packageExcludes: string;
  recommendedGear: string;
  difficultyLevel: string;
  price: number;
}