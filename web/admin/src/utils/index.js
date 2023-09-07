import { isRef } from "vue";
import { notify } from "notiwind";

export function getCookie(cname) {
  let name = cname + "=";
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(";");
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

export function setCookie(cname, cvalue, days) {
  let date = new Date();
  date.setDate(date.getDate() + days);
  let value = cvalue + (days == null ? "" : "; expires=" + date.toUTCString());
  document.cookie = cname + "=" + value;
}

export function getUser() {
  let obj = getCookie("loginUser");
  if (obj) {
    return JSON.parse(obj);
  }
}

export function setUser(user) {
  if (!user) {
    return;
  }

  let value = JSON.stringify(user);
  setCookie("loginUser", value, 1);
}

export function delUser() {
  setCookie("loginUser", "", -1);
}

export function updateItem(list, row, newVal) {
  if (isRef(list)) {
    list = list.value;
  }
  let index = getItemIndex(list, row);
  list[index] = newVal;
}

export function deleteItem(list, row) {
  if (isRef(list)) {
    list = list.value;
  }
  let index = getItemIndex(list, row);
  list.splice(index, 1);
}

export function getItemIndex(list, row) {
  for (let i in list) {
    if (deepEqual(list[i], row)) {
      return i;
    }
  }
}

export function costFormat(cost) {
  return Number(cost) ? (Number(cost) / 100).toFixed(2) : "0.00";
}

export function costStripe(cost) {
  return Number(cost) * 100;
}

export function formatDate(timestamp) {
  const monthNames = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
  const date = new Date(timestamp * 1000);
  return `${date.getDate()} ${
    monthNames[date.getMonth()]
  } ${date.getFullYear()}, ${date.toLocaleTimeString()}`;
}

function deepEqual(object1, object2) {
  if (object1 == object2) {
    return true;
  }

  const keys1 = Object.keys(object1);
  const keys2 = Object.keys(object2);

  if (keys1.length !== keys2.length) {
    return false;
  }

  for (let index = 0; index < keys1.length; index++) {
    const val1 = object1[keys1[index]];
    const val2 = object2[keys2[index]];
    const areObjects = isObject(val1) && isObject(val2);
    if (
      (areObjects && !deepEqual(val1, val2)) ||
      (!areObjects && val1 !== val2)
    ) {
      return false;
    }
  }

  return true;
}

function isObject(object) {
  return object != null && typeof object === "object";
}
