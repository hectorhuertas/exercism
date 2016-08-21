var Person = function(year){
  this.year = year;
};

Person.prototype.isLeap = function(){
  if (this.divByFour() && !this.divBy100()) { return true }
  if (this.divByFour() &&  this.divBy400()) { return true }
  return false
};

Person.prototype.divByFour = function(){
  return this.year % 4 === 0
}
Person.prototype.divBy100 = function(){
  return this.year % 100 === 0
}
Person.prototype.divBy400 = function(){
  return this.year % 400 === 0
}

module.exports = Person;
