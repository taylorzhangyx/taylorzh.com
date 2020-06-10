import { Component, OnInit } from '@angular/core';
import { map } from 'rxjs/operators';

import { environment } from '@env/environment';
import { HttpClient } from '@angular/common/http';
import { Logger } from '@app/@core';

const logger = new Logger('about');

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.scss'],
})
export class AboutComponent implements OnInit {
  version: string | null = environment.version;
  v_count: number = 0;
  constructor(private httpClient: HttpClient) {}

  ngOnInit() {
    this.httpClient.get('/count', { responseType: 'json' }).subscribe((value: { count: number; message: string }) => {
      console.log(value);
      logger.info('reqested gin -> ' + value);
      this.v_count = value.count;
    });
  }
}
