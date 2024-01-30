import { UserTransferExt } from "./Transfer";
import { FoundrisingTransfer} from "../foundrising/Transfer";
const baseUrl = 'http://localhost:8080/api/v1';
const url = `${baseUrl}/users`;

function translateStatusToErrorMessage(status: number) {
  switch (status) {
    case 401:
      return 'Please login again.';
    case 403:
      return 'You do not have permission to view the User(s).';
    default:
      return 'There was an error retrieving the User(s). Please try again.';
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


function convertToUserModel(item: any): UserTransferExt {
  //console.log(item)
  return new UserTransferExt(item);
}
function convertToSingleUserModel(item: any): UserTransferExt {
  //console.log(item)
  return new UserTransferExt(item.User);
}

const UserAPI = {
  
  find(id: number){
    return (
      fetch(`${url}/${id}`)
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleUserModel)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving one User. Please try again.'
        );
      })
    )
  },
  findExtLogin(login: string, token: string){
    console.log("FINDING EXT WITH TOKEN: ", token);
    return (
      fetch(`${url}/login/${login}`, {
        method: "GET",
        headers: {
          'Authorization': `${token}`
        }
      })
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleUserModel) //мб тут не Single
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving extended User by login. Please try again.'
        );
      })
    )
  },
  donate(token: string, id: string, donBody: string){ // may be id: number
    return (
      fetch(`${url}/${id}/donate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        },
        body: donBody
        })
        .then((response) => {
          console.log(response);
          return response;
        })
        .then(checkStatus)
        .then(parseJSON)
        .catch((error: TypeError) => {
          console.log('log client error ' + error);
          throw new Error(
            'There was an error donating From User to Foundrising. Please try again.'
          );
        })
    )
  },
  replenish(token: string, id: string, replBody: string){ // may be id: number
    return (
      fetch(`${url}/${id}/replenish`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        },
        body: replBody
        })
        .then((response) => {
          console.log(response);
          return response;
        })
        .then(checkStatus)
        .then(parseJSON)
        .catch((error: TypeError) => {
          console.log('log client error ' + error);
          throw new Error(
            'There was an error replenishing User Balance. Please try again.'
          );
        })
    )
  },
  delete(token: string, id: string){ // may be id: number
    return (
      fetch(`${url}/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        }
        })
        .then((response) => {
          console.log(response);
          return response;
        })
        .then(checkStatus)
        .then(parseJSON)
        .catch((error: TypeError) => {
          console.log('log client error ' + error);
          throw new Error(
            'There was an error deleting User. Please try again.'
          );
        })
    )
  },
  update(token: string, id: string, updBody: string){ // may be id: number
    return (
      fetch(`${url}/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        },
        body: updBody
        })
        .then((response) => {
          console.log(response);
          return response;
        })
        .then(checkStatus)
        .then(parseJSON)
        .catch((error: TypeError) => {
          console.log('log client error ' + error);
          throw new Error(
            'There was an error updating User. Please try again.'
          );
        })
    )
  }
};
export { UserAPI };
//1.  
//    Update,
//    Delete,
//    Donate,
//    Replenish
//2.
//    Foundations -> foundation/id -> foundrising -> foundrisings/id 