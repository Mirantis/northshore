import {Pipe, PipeTransform} from '@angular/core';

// http://stackoverflow.com/a/37479557
@Pipe({ name: 'iterateMap' })
export class IterateMapPipe implements PipeTransform {
  transform(map: {}, args: any[] = null): any {
    if (!map)
      return null;
    return Object.keys(map)
      .map((key) => ({ 'key': key, 'value': map[key] }));
  }
}


@Pipe({ name: 'keys' })
export class KeysPipe implements PipeTransform {
  transform(map: {}): string[] {
    return Object.keys(map)
  }
}


class SumIfValuePipeOptions {
  filter: string[] = [];
  hideZero: boolean = false;
}

@Pipe({ name: 'sumIfValue' })
export class SumIfValuePipe implements PipeTransform {
  transform(map: {}, options: SumIfValuePipeOptions): number {
    let sum: number = 0;
    // ES7 Array.prototype.includes() missing in TypeScript #2340
    for (let key in map) {
      if (options.filter.indexOf(map[key]) > -1) {
        sum++;
      }
    }
    return (sum === 0 && options.hideZero) ? null : sum;
  }
}
