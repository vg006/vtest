export interface Response {
  StatusCode: number;
  Headers: Header[];
}

export interface Header {
  Key: string;
  Value: string;
}
