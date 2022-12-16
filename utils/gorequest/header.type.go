package gorequest

// HeaderIsImg 判断是否为图片
func (r *Response) HeaderIsImg() bool {
	if r.ResponseHeader.Get("Content-Type") == "image/jpeg" || r.ResponseHeader.Get("Content-Type") == "image/png" || r.ResponseHeader.Get("Content-Type") == "image/jpg" {
		return true
	}
	return false
}

// HeaderIsJpeg 判断是否为jpeg图片
func (r *Response) HeaderIsJpeg() bool {
	if r.ResponseHeader.Get("Content-Type") == "image/jpeg" {
		return true
	}
	return false
}

// HeaderIsPng 判断是否为Png图片
func (r *Response) HeaderIsPng() bool {
	if r.ResponseHeader.Get("Content-Type") == "image/png" {
		return true
	}
	return false
}

// HeaderIsJpg 判断是否为Jpg图片
func (r *Response) HeaderIsJpg() bool {
	if r.ResponseHeader.Get("Content-Type") == "image/jpg" {
		return true
	}
	return false
}

// HeaderJson 判断是否为Json数据
func (r *Response) HeaderJson() bool {
	if r.ResponseHeader.Get("Content-Type") == "application/json" {
		return true
	}
	return false
}

// HeaderHtml 判断是否为Html
func (r *Response) HeaderHtml() bool {
	if r.ResponseHeader.Get("Content-Type") == "text/html" || r.ResponseHeader.Get("Content-Type") == "application/xhtml+xml" {
		return true
	}
	return false
}

// HeaderTextHtml 判断是否为Html
func (r *Response) HeaderTextHtml() bool {
	if r.ResponseHeader.Get("Content-Type") == "text/html" {
		return true
	}
	return false
}

// HeaderXHtml 判断是否为Html
func (r *Response) HeaderXHtml() bool {
	if r.ResponseHeader.Get("Content-Type") == "application/xhtml+xml" {
		return true
	}
	return false
}
