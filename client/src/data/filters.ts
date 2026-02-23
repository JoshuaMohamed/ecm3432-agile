export const filterOptions = ["Area", "District", "Sector"] as const;

export type FilterOption = (typeof filterOptions)[number];

export const filterParam = (filter: FilterOption) => filter.toLowerCase();
