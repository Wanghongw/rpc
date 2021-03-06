function getVal(obj) {
	if(document.getElementById(obj) != null)
		return document.getElementById(obj).value;
	else return 0;
}

function setVal(obj, val) {
	if(document.getElementById(obj) != null)
		document.getElementById(obj).value = val;

}

function Convert2Dec() {
	var deg = Math.abs(getVal('deg'));
	var min = Math.abs(getVal('min'));
	var sec = Math.abs(getVal('sec'));
	var deci = deg * 1 + (sec * 1 + min * 60) / 3600;
	setVal("deg2", deci);
}

function Convert2Deg() {
	var deci = Math.abs(getVal('deg2'));
	var deci2 = deci + '';

	if(deci2.indexOf('.') == -1) {
		setVal("deg", deci);
		return false;
	}
	deci = deci2.split(".");
	setVal("deg", deci[0]);

	//
	deci[1] = "0." + deci[1];
	var min_sec = deci[1] * 3600;
	var min = Math.floor(min_sec / 60);
	var sec = (min_sec - (min * 60));

	setVal("min", min);

	setVal("sec", sec);

}

function hide(m) {
	document.getElementById(m).style.display = "none";
	return true;
}

function show(m) {
	document.getElementById(m).style.display = "";
	return true;
}
//private const double EARTH_RADIUS = 6378.137;
function rad(d) {
	return d * Math.PI / 180.0;
}

function GetDistance(lat1, lng1, lat2, lng2) {
	hide("warning");
	if((Math.abs(lat1) > 90) || (Math.abs(lat2) > 90)) {
		document.getElementById("warning").innerHTML = ("兄台，这哪里是纬度啊？分明是想忽悠我嘛");
		show("warning");
		return "耍我？拒绝计算！";
	} else {
		hide("warning");
	}
	if((Math.abs(lng1) > 180) || (Math.abs(lng2) > 180)) {

		show("warning");
		document.getElementById("warning").innerHTML = ("兄台，这哪里是经度啊？分明是想忽悠我嘛");
		return "耍我？拒绝计算！";
	} else {
		hide("warning");
	}
	var radLat1 = rad(lat1);
	var radLat2 = rad(lat2);
	var a = radLat1 - radLat2;
	var b = rad(lng1) - rad(lng2);
	var s = 2 * Math.asin(Math.sqrt(Math.pow(Math.sin(a / 2), 2) +
		Math.cos(radLat1) * Math.cos(radLat2) * Math.pow(Math.sin(b / 2), 2)));
	s = s * 6378.137; // EARTH_RADIUS;
	s = Math.round(s * 10000) / 10000;
	return s;
}

function calDis() {
	var lat1 = document.getElementById("lat1").value * 1;
	var lat2 = document.getElementById("lat2").value * 1;
	var lng1 = document.getElementById("lng1").value * 1;
	var lng2 = document.getElementById("lng2").value * 1;
	var dis = GetDistance(lat1, lng1, lat2, lng2);
	document.getElementById("distance").value = dis;

}