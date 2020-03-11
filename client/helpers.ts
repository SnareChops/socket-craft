export function getQueryParams(): { [key: string]: string } {
  if (!!location.search) {
    return location.search
      .slice(1)
      .split('&')
      .reduce((result, x) => {
        const [key, value] = x.split('=');
        result[key] = value;
        return result;
      }, {});
  }
  return {};
}
