module PersonalHelper
  def personal_bar

    links = []

    if current_user
      if current_user.is_admin?
        links << {href: site_path, title: 'site.index.title'}
        links << {href: site_users_path, title: 'site.users.title'}
        links << {href: cms_tags_path, title: 'cms.tags.index.title'}
        links << {href: cms_comments_path, title: 'cms.comments.index.title'}
        links << {href: questionnaire_reports_path, title: 'questionnaire.reports.index.title'}
      end
      links << {href: edit_user_registration_path, title: t('devise.registrations.edit.title')}
      links << {href: attachments_path, title: t('attachments.index.title')}
      links << {href: personal_logs_path, title: 'personal.logs.title'}
    end
    links
  end
end