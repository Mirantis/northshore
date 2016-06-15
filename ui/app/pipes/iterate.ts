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

// http://stackoverflow.com/a/35536052
@Pipe({ name: 'keys' })
export class KeysPipe implements PipeTransform {
  transform(map: {}, args: any[] = null): any {
    if (!map)
      return null;
    let keys: string[] = [];
    for (let key in map) {
      keys.push(key);
    }
    return keys;
  }
}
