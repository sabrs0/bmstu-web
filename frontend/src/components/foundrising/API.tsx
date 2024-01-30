import { FoundrisingTransfer } from "./Transfer";
const baseUrl = 'http://localhost:8080/api/v1';
const url = `${baseUrl}/foundrisings`;

function translateStatusToErrorMessage(status: number) {
  switch (status) {
    case 401:
      return 'Please login again.';
    case 403:
      return 'You do not have permission to view the Foundrising(s).';
    default:
      return 'There was an error retrieving the Foundrising(s). Please try again.';
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

function convertToFoundrisingModels(data: any): FoundrisingTransfer[] {
  console.log("FOUNDRISINGS ARE: ", data)
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

const FoundrisingAPI = {
  get(page = 1, limit = 20) {
    return fetch(`${url}`)
      .then(delay(600))
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToFoundrisingModels)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving the Foundrisings. Please try again.'
        );
      });
  },
  find(id: number){
    return (
      fetch(`${url}/${id}`)
      .then(checkStatus)
      .then(parseJSON)
      .then(convertToSingleFoundrisingModel)
      .catch((error: TypeError) => {
        console.log('log client error ' + error);
        throw new Error(
          'There was an error retrieving one Foundrising. Please try again.'
        );
      })
    )
  },
  add(token: string, addBody: string){
    return (
      // Отправка POST запроса на сервер с FoundrisingData
      fetch(`${url}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${token}`
        },
        body: addBody
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
            'There was an error adding Foundrising. Please try again.'
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
            'There was an error updating Foundrising. Please try again.'
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
            'There was an error deleting Foundrising. Please try again.'
          );
        })
    )
  }
};

export { FoundrisingAPI };