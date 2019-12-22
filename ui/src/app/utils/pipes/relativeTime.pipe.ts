import { Pipe, PipeTransform } from '@angular/core';

@Pipe({name: 'relativeTime'})
export class RelativeTimePipe implements PipeTransform {
  transform(value: Date): string {
    const un = (new Date(value)).getTime();
    const minute = 1000 * 60;
    const hour = minute * 60;
    const day = hour * 24;
    const halfamonth = day * 15;
    const month = day * 30;
    const now = new Date().getTime();
    const diffValue = now - un;
    if (diffValue < 0) {
      return;
    }
    const monthC = diffValue / month;
    const weekC = diffValue / (7 * day);
    const dayC = diffValue / day;
    const hourC = diffValue / hour;
    const minC = diffValue / minute;
    if (monthC >= 1) {
      return '' + Math.floor(monthC) + '月前';
    } else if (weekC >= 1) {
      return '' + Math.floor(weekC) + '周前';
    } else if (dayC >= 1) {
      return  '' + Math.floor(dayC) + '天前';
    } else if (hourC >= 1) {
      return '' + Math.floor(hourC) + '小时前';
    } else if (minC >= 1) {
      return '' + Math.floor(minC) + '分钟前';
    } else {
      return '刚刚';
    }
    return value.toLocaleDateString();
  }
}
