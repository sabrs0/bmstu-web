export class FoundationTransfer {
    id: number | undefined;
    name: string = '';
    cur_foudrising_amount: number = 0;
    closed_foundrising_amount: number = 0;
    volunteer_amount: number = 0;
    country: string = '';
    get isNew(): boolean {
      return this.id === undefined;
    }
  
    constructor(initializer?: any) {
      if (!initializer) return;
      if (initializer.id){
        this.id = initializer.id;
      }
      this.name = initializer.name;
      this.cur_foudrising_amount = initializer.cur_foudrising_amount;
      this.closed_foundrising_amount = initializer.closed_foundrising_amount;
      this.volunteer_amount = initializer.volunteer_amount;
      this.country = initializer.country;
    }
  }
  export class FoundationExt {
    id: number | undefined;
    name: string = '';
    cur_foudrising_amount: number = 0;
    closed_foundrising_amount: number = 0;
    volunteer_amount: number = 0;
    country: string = '';
    fund_balance: number = 0.0;
    income_history: number = 0.0;
    outcome_history: number = 0.0;
    login: string = '';
    password: string = '';
    get isNew(): boolean {
      return this.id === undefined;
    }
  
    constructor(initializer?: any) {
      if (!initializer) return;
      if (initializer.id){
        this.id = initializer.id;
      }
      this.name = initializer.name;
      this.cur_foudrising_amount = initializer.cur_foudrising_amount;
      this.closed_foundrising_amount = initializer.closed_foundrising_amount;
      this.volunteer_amount = initializer.volunteer_amount;
      this.country = initializer.country;
      this.fund_balance = initializer.fund_balance;
      this.income_history = initializer.income_history;
      this.outcome_history = initializer.outcome_history;
      this.login = initializer.login;
      this.password = initializer.password;
    }
  }

export class FoundationDonate {
  foundrising_id: string = '';
  sum_of_money: string = '';
  comment: string = '';

  constructor(initializer?: any){
    this.foundrising_id = initializer.foundrising_id;
    this.sum_of_money = initializer.sum_of_money;
    this.comment = initializer.comment;
  }
}
export class FoundationReplenish {
  sum_of_money: string = '';
  
  constructor(initializer?: any){
    this.sum_of_money = initializer.sum_of_money;
  }
}
export class FoundationPut {
  name: string = '';
  login: string = '';
  password: string = '';
  country: string = '';
  
  constructor(initializer?: any){
    this.name = initializer.name;
    this.login = initializer.login;
    this.password = initializer.password;
    this.country = initializer.country;
  }
}