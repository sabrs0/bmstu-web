export class LoginTransfer {
    login: string = '';
    password: string = '';
    role: string = '';
  
    constructor(initializer?: any) {
      if (!initializer) return;
      this.login = initializer.login;
      this.password = initializer.password;
      this.role = initializer.role;
    }
  }
export class LoginToken {
  token: string = '';
  constructor(initializer?: any) {
    if (!initializer) return;
    this.token = initializer.token;
  }
}