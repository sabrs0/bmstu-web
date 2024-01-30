export class FoundrisingTransfer {
    id: number | undefined;
    found_id: number = -1;
    description: string = '';
    required_sum: number = 0;
    current_sum: number = 0;
    philantrops_amount: number = 0;
    creation_date: string = '';
    closing_date: string  | undefined;
    get isNew(): boolean {
      return this.id === undefined;
    }
  
    constructor(initializer?: any) {
      if (!initializer) return;
      if (initializer.id){
        this.id = initializer.id;
      }
      this.found_id = initializer.found_id;
      this.description = initializer.description;
      this.required_sum = initializer.required_sum;
      this.current_sum = initializer.current_sum;
      this.philantrops_amount = initializer.philantrops_amount;
      this.creation_date = initializer.creation_date;
      if (initializer.closing_date.Valid){
        this.closing_date = initializer.closing_date.String;
      }
      
    }
  }
  export class FoundrisingAdd {
    found_id: number = -1;
    description: string = '';
    required_sum: string = '';
    creation_date: string = '';
  
    constructor(initializer?: any) {
      if (!initializer) return;
      
      this.found_id = initializer.found_id;
      this.description = initializer.description;
      this.required_sum = initializer.required_sum;
      this.creation_date = initializer.creation_date;
    }
  }
  export class FoundrisingPut {
    description: string = '';
    required_sum: string = '';
  
    constructor(initializer?: any) {
      if (!initializer) return;
      
      this.description = initializer.description;
      this.required_sum = initializer.required_sum;
    }
  }