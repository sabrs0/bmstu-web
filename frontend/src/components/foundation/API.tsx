import { FoundationTransfer } from "./Transfer";
import { FoundationExt } from "./Transfer";
import { FoundrisingTransfer} from "../foundrising/Transfer";
const baseUrl = 'http://localhost:8080/api/v1';
const url = `${baseUrl}/foundations`;

function translateStatusToErrorMessage(status: number) {
  switch (status) {
    case 401:
      return 'Please login again.';
    case 403:
      return 'You do not have permission to view the Foundation(s).';
    default:
      return 'There was an error retrieving the Foundation(s). Please try again.';
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

function convertToFoundationModels(data: any): FoundationTransfer[] {
  //console.log(data)
  let Foundations: FoundationTransfer[] = data.foundations.map(convertToFoundationModel);
  console.log(Foundations);
  return Foundations;
}

function convertToFoundationModel(item: any): FoundationTransfer {
  //console.log(item)
  return new FoundationTransfer(item);
}
function convertToSingleFoundationModel(item: any): FoundationTransfer {
  //console.log(item)
  return new FoundationTransfer(item.Foundation);
}

function convertToExtendedModel(item: any): FoundationExt {
  //console.log(item)
  return new FoundationExt(item);
}
function convertToSingleExtendedModel(item: any): FoundationExt {
  //console.log(item)
  return new FoundationExt(item.Foundation);
}


function convertToFoundrisingModels(data: any): FoundrisingTransfer[] {
  //console.log(data)
  let Foundrisings: FoundrisingTransfer[] = data.foundrisings.map(convertToFoundrisingModel);
  console.log(Foundrisings);
  return Foundrisings;
}

function convertToFoundrisingModel(item: any): FoundrisingTransfer {
  //console.log(item)
  return new FoundrisingTransfer(item);
}
function convertToSingleFoundrisingModel(item: any): FoundrisingTransfer {
  //console.log(item)
  return new FoundrisingTransfer(item.Foundrising);
}
const FoundationAPI = {
  get(page = 1, limit = 10) {
    return fetch(`${url}`) //?_page=${page}&_limit=${limit}
      .then(delay(600))
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToFoundationModels)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving the Foundations. Please try again.'
        );
      });
  },
  find(id: number){
    return (
      fetch(`${url}/${id}`)
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleFoundationModel)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving one Foundation. Please try again.'
        );
      })
    )
  },
  findExt(id: number, token: string){
    console.log("FINDING EXT WITH TOKEN: ", token);
    return (
      fetch(`${url}/extended/${id}`, {
        method: "GET",
        headers: {
          'Authorization': `${token}`
        }
      })
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleExtendedModel) //мб тут не Single
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving one Foundation. Please try again.'
        );
      })
    )
  },
  findExtLogin(login: string, token: string){
    console.log("FINDING EXT WITH TOKEN: ", token);
    return (
      fetch(`${url}/extended_login/${login}`, {
        method: "GET",
        headers: {
          'Authorization': `${token}`
        }
      })
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleExtendedModel) //мб тут не Single
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving extended Foundation by login. Please try again.'
        );
      })
    )
  },
  findFrisingsByID(id: number, page = 1, limit = 10){
    return (
      fetch(`${url}/${id}/foundrisings`) //?_page=${page}&_limit=${limit}
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToFoundrisingModels)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving Foundation foundrisings. Please try again.'
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
            'There was an error donating From Foundation to Foundrising. Please try again.'
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
            'There was an error replenishing Foundation Balance. Please try again.'
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
            'There was an error deleting Foundation. Please try again.'
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
            'There was an error updating Foundation. Please try again.'
          );
        })
    )
  }
};

export { FoundationAPI };