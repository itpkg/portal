module PersonalHelper
  def personal_bar

    links = []

    if current_user
      if current_user.is_admin?
        links << {href: site_path, title: 'site.index.title'}
        links << {href: site_users_path, title: 'site.users.title'}
      end
      links << {href: edit_user_registration_path, title: t('devise.registrations.edit.title')}
      links << {href: personal_logs_path, title: 'personal.logs.title'}
    end
    links
  end
end