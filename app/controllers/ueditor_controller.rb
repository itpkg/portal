# http://fex-team.github.io/ueditor/#dev-request_specification
# http://fex-team.github.io/ueditor/#dev-serverparam


class UeditorController < ApplicationController
  before_action :authenticate_user!
  layout false

  def index

    case __action
      when :config
        rv = __config
      when :upload_image
        rv = __store __field(:imageFieldName)
      when :upload_scrawl
        file = Tempfile.new %w(scrawl .png), encoding: 'ascii-8bit'
        begin
          file.write Base64.strict_decode64(__field(:scrawlFieldName))
          file.flush
          rv = __store_t file.path, file, 'image/png'
        ensure
          file.close
          file.unlink
        end
      when :upload_video
        rv = __store __field(:videoFieldName)
      when :upload_file
        rv = __store __field(:fileFieldName)
      when :list_image
        rv = __list __config[:imageManagerAllowFiles]
      when :list_file
        rv = __list __config[:fileManagerAllowFiles]
      when :catch_image
        rv = __remote_image(:catcherFieldName)

      else
        rv = {stat: t('ueditor.bad_request')}
    end

    __render rv
  end

  private

  def __remote_image(urls)
    files = urls.map do |url|
      ret = {}
      res = Net::HTTP.get_response URI(url)
      file = Tempfile.new ['remote', File.extname(url)], encoding: 'ascii-8bit'
      begin
        if res.is_a?(Net::HTTPSuccess)
          if res.content_type.include?('image')
            file.write res.body
            file.flush
            ret = __store_t url, file, res.content_type
            ret[:source] = url
          else
            ret[:state] = t 'ueditor.not_picture'
          end
        else
          ret[:state] = t 'ueditor.http_error'
        end
      ensure
        file.close
        file.unlink
      end
      ret
    end
    {
        state: (files.empty? ? t('ueditor.error') : t('ueditor.success')),
        list: files
    }
  end

  def __list(types)
    files = Attachment.order(id: :desc).where(user_id: current_user.id).select { |a| types.include? File.extname(a.avatar_identifier) }.map { |a| {url: a.avatar_url} }
    {
        state: t('ueditor.success'),
        list: files,
        start: 0,
        total: files.size
    }
  end

  def __store_t(n, f, t)
    a = Attachment.new user: current_user
    a.title = n
    a.size = f.size
    a.content_type = t
    a.avatar = f
    __save a
  end

  def __store(f)
    a = Attachment.new user: current_user
    a.read! f
    __save a
  end

  def __save(a)
    if a.save
      {
          state: t('ueditor.success'),
          url: a.avatar_url,
          title: a.title,
          original: a.title,
      }
    else
      {
          state: t('ueditor.error')
      }
    end
  end

  def __field(k)
    params.fetch __config.fetch(k)
  end

  def __render(rv)
    cb = params[:callback]
    if cb
      if cb =~ /^[\w_]+$/
        render plain: "<script>#{cb}(#{rv})</script"
      else
        render json: {state: t('ueditor.bad_callback')}
      end
    else
      render json: rv
    end
  end

  def __action
    CGI.parse(URI.parse(request.url).query)['action'].first.to_sym
  end

  def __config
# 前后端通信相关的配置,注释只允许使用多行方式
    {
        # 上传图片配置项
        imageActionName: 'upload_image', # 执行上传图片的action名称
        imageFieldName: :up_file, # 提交的图片表单名称
        imageMaxSize: 2048000, # 上传大小限制，单位B
        imageAllowFiles: %w(.png .jpg .jpeg .gif .bmp), # 上传图片格式显示
        imageCompressEnable: true, # 是否压缩图片,默认是true
        imageCompressBorder: 1600, # 图片压缩最长边限制
        imageInsertAlign: 'none', # 插入的图片浮动方式
        imageUrlPrefix: '', # 图片访问路径前缀
        imagePathFormat: '/ueditor/ror/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        # {filename} 会替换成原文件名,配置这项需要注意中文乱码问题
        # {rand:6} 会替换成随机数,后面的数字是随机数的位数
        # {time} 会替换成时间戳
        # {yyyy} 会替换成四位年份
        # {yy} 会替换成两位年份
        # {mm} 会替换成两位月份
        # {dd} 会替换成两位日期
        # {hh} 会替换成两位小时
        # {ii} 会替换成两位分钟
        # {ss} 会替换成两位秒
        # 非法字符 \ : * ? ' < > |
        # 具请体看线上文档: fex.baidu.com/ueditor/#use-format_upload_filename

        # 涂鸦图片上传配置项
        scrawlActionName: 'upload_scrawl', # 执行上传涂鸦的action名称
        scrawlFieldName: :up_file, # 提交的图片表单名称
        scrawlPathFormat: '/ueditor/ror/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        scrawlMaxSize: 2048000, # 上传大小限制，单位B
        scrawlUrlPrefix: '', # 图片访问路径前缀
        scrawlInsertAlign: 'none',

        # 截图工具上传
        snapscreenActionName: 'upload_image', # 执行上传截图的action名称
        snapscreenPathFormat: '/ueditor/ror/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        snapscreenUrlPrefix: '', # 图片访问路径前缀
        snapscreenInsertAlign: 'none', # 插入的图片浮动方式

        # 抓取远程图片配置
        catcherLocalDomain: %w(127.0.0.1 localhost img.baidu.com),
        catcherActionName: 'catch_image', # 执行抓取远程图片的action名称
        catcherFieldName: :source, # 提交的图片列表表单名称
        catcherPathFormat: '/ueditor/ror/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        catcherUrlPrefix: '', # 图片访问路径前缀
        catcherMaxSize: 2048000, # 上传大小限制，单位B
        catcherAllowFiles: %w(.png .jpg .jpeg .gif .bmp), # 抓取图片格式显示

        # 上传视频配置
        videoActionName: 'upload_video', # 执行上传视频的action名称
        videoFieldName: :up_file, # 提交的视频表单名称
        videoPathFormat: '/ueditor/ror/upload/video/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        videoUrlPrefix: '', # 视频访问路径前缀
        videoMaxSize: 102400000, # 上传大小限制，单位B，默认100MB
        videoAllowFiles: %w(
            .flv .swf .mkv .avi .rm .rmvb .mpeg .mpg
            .ogg .ogv .mov .wmv .mp4 .webm .mp3 .wav .mid), # 上传视频格式显示

        # 上传文件配置
        fileActionName: 'upload_file', # controller里,执行上传视频的action名称
        fileFieldName: :up_file, # 提交的文件表单名称
        filePathFormat: '/ueditor/ror/upload/file/{yyyy}{mm}{dd}/{time}{rand:6}', # 上传保存路径,可以自定义保存路径和文件名格式
        fileUrlPrefix: '', # 文件访问路径前缀
        fileMaxSize: 51200000, # 上传大小限制，单位B，默认50MB
        fileAllowFiles: %w(
            .png .jpg .jpeg .gif .bmp
            .flv .swf .mkv .avi .rm .rmvb .mpeg .mpg
            .ogg .ogv .mov .wmv .mp4 .webm .mp3 .wav .mid
            .rar .zip .tar .gz .7z .bz2 .cab .iso
            .doc .docx .xls .xlsx .ppt .pptx .pdf .txt .md .xml
        ), # 上传文件格式显示

        # 列出指定目录下的图片
        imageManagerActionName: 'list_image', # 执行图片管理的action名称
        imageManagerListPath: '/ueditor/ror/upload/image/', # 指定要列出图片的目录
        imageManagerListSize: 20, # 每次列出文件数量
        imageManagerUrlPrefix: '', # 图片访问路径前缀
        imageManagerInsertAlign: 'none', # 插入的图片浮动方式
        imageManagerAllowFiles: %w(.png .jpg .jpeg .gif .bmp), # 列出的文件类型

        # 列出指定目录下的文件
        fileManagerActionName: 'list_file', # 执行文件管理的action名称
        fileManagerListPath: '/ueditor/ror/upload/file/', # 指定要列出文件的目录
        fileManagerUrlPrefix: '', # 文件访问路径前缀
        fileManagerListSize: 20, # 每次列出文件数量
        fileManagerAllowFiles: %w(
            .png .jpg .jpeg .gif .bmp
            .flv .swf .mkv .avi .rm .rmvb .mpeg .mpg
            .ogg .ogv .mov .wmv .mp4 .webm .mp3 .wav .mid
            .rar .zip .tar .gz .7z .bz2 .cab .iso
            .doc .docx .xls .xlsx .ppt .pptx .pdf .txt .md .xml
        ) # 列出的文件类型

    }
  end


end