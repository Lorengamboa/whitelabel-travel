export const replaceParams = (url: string, params: Record<string, string>) => {
  const entries = Object.entries(params).reduce((acc, [key, value]) => {
    return acc.replace(`:${key}`, value);
  }, url);


  // trim the unreplaces params
  return entries.replace(/\/:\w+/g, '');
}