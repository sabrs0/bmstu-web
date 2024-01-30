import { LoginToken, LoginTransfer } from "./LoginTransfer";
const baseUrl = 'http://localhost:8080/api/v1';
const url = `${baseUrl}/login`;

function translateStatusToErrorMessage(status: number) {
  switch (status) {
    case 401:
      return 'Please login again.';
    case 403:
      return 'You do not have permission to view the Login(s).';
    default:
      return 'There was an error retrieving the Login(s). Please try again.';
  }
}

function checkStatus(response: any) {
  if (response.ok) {
    return response;
  } else {
    const httpErrorInfo = {
      status: response.status,
      statusText: response.statusText,
      url: response.url,
    };
    console.log(`log server http error: ${JSON.stringify(httpErrorInfo)}`);

    let errorMessage = translateStatusToErrorMessage(httpErrorInfo.status);
    throw new Error(errorMessage);
  }
}

function parseJSON(response: Response) {
  return response.json();
}

// eslint-disable-next-line
function delay(ms: number) {
  return function (x: any): Promise<any> {
    return new Promise((resolve) => setTimeout(() => resolve(x), ms));
  };
}

function convertToToken(item: any): LoginToken {
  console.log("CONVERING... :",item)
  return new LoginToken(item);
}
function convertToSingleToken(item: any): LoginToken {
  //console.log(item)
  return new LoginToken(item.token);
}

const LoginAPI = {
  
  login(login: string, password: string, role: string){
    return (
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({login, password, role}),
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then((response) => {
          console.log(response);
          return response;
        })
        .then(checkStatus)
        .then(parseJSON)
        //.then((response ) => {return convertToToken(response)})
        //.then((response) => console.log('Success:', JSON.stringify(response)))
        .catch((error: TypeError) => {
          console.log('log client error ' + error);
          throw new Error(
            'There was an error retrieving one Login. Please try again.'
          );
        }) 
    );
  }
};

export { LoginAPI };