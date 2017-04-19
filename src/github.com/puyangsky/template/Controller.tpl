package com.pku.osa.web.controller;


import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@RequestMapping(value="/")
{{.AccessMode}} {{.ClassType}} {{.ClassName}}Controller {


	@RequestMapping(value="/list.html", method=RequestMethod.GET, produces="application/json;charset=UTF-8")
	public index(HttpServletRequest request){

		return adbanners;
	}
}
