function getWeekday(dateString) {
    const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    const date = new Date(dateString);
    return daysOfWeek[date.getUTCDay()];
  }
  
  console.log(getWeekday("1912-11-11"));
  