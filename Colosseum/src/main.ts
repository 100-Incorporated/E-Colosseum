import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';

const loginpage = document.querySelector('.loginpage');
const bootstrapPromise = platformBrowserDynamic().bootstrapModule(AppModule);

bootstrapPromise.then(success => console.log(`Bootstrap success`))
  .catch(err => console.error(err));

