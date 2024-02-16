export class UserTransferExt {
    id: number | undefined;
    login: string = '';
    password: string = '';
    balance: number = 0.00;
    charity_sum: number = 0.00;
    closed_fing_amount: number = 0;
    get isNew(): boolean {
      return this.id === undefined;
    }
  
    constructor(initializer?: any) {
      if (!initializer) return;
      if (initializer.id){
        this.id = initializer.id;
      }
      this.login = initializer.login;
      this.password = initializer.password;
      this.balance = initializer.balance;
      this.charity_sum = initializer.charity_sum;
      this.closed_fing_amount = initializer.closed_fing_amount;
    }
  }

  export class UserDonate {
    
    entity_type: boolean = false;
    entity_id: string = '';
    sum_of_money: string = '';
    comment: string = '';

    constructor(initializer?: any){
      if (!initializer) return;
      this.entity_type = initializer.entity_type;
      this.entity_id = initializer.entity_id;
      this.sum_of_money = initializer.sum_of_money;
      this.comment = initializer.comment;
    }
  }
export class UserReplenish {
  sum_of_money: string = '';
  
  constructor(initializer?: any){
    this.sum_of_money = initializer.sum_of_money;
  }
}
export class UserPut {
  login: string = '';
  password: string = '';
  
  constructor(initializer?: any){
    this.login = initializer.login;
    this.password = initializer.password;
  }
}