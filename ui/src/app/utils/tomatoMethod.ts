export default function transformTime(value: Date): string {
  const un = (new Date(value)).getTime();
  const minute = 1000 * 60;
  const second = 1000;
  const now = new Date().getTime();
  const diffValue = now - un;
  if (diffValue < 0) {
    return '25:00';
  }
  const minC = Math.floor(diffValue / minute);
  if (minC >= 25) {
    return '已完成';
  }
  const secC = 59 - Math.floor((diffValue - minC * minute) / second);

  return ('00' + (24 - minC)).slice(-2) + ':' + ('00' + secC).slice(-2);
}
