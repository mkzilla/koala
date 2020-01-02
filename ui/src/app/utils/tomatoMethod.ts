export default function transformTime(value: Date): string {
  const un = (new Date(value)).getTime();
  const minute = 1000 * 60;
  const second = 1000;
  const now = new Date().getTime();
  const diffValue = now - un;
  if (diffValue < 0) {
    return;
  }
  const minC = Math.floor(diffValue / minute);
  if (minC >= 25) {
    this.rain.nativeElement.pause();
    return '已完成';
  }
  const secC = Math.floor((diffValue - minC * minute) / second);
  return ('00' + minC).slice(-2) + ':' + ('00' + secC).slice(-2);
}
